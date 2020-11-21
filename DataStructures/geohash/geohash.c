#include "geohash.h"
#include "sds.h"
#include <stdio.h>

static inline uint64_t interleave64(uint32_t xlo, uint32_t ylo)
{
	static const uint64_t B[] = {0x5555555555555555ULL, 0x3333333333333333ULL,
								 0x0F0F0F0F0F0F0F0FULL, 0x00FF00FF00FF00FFULL,
								 0x0000FFFF0000FFFFULL};
	static const unsigned int S[] = {1, 2, 4, 8, 16};

	uint64_t x = xlo;
	uint64_t y = ylo;

	x = (x | (x << S[4])) & B[4];
	y = (y | (y << S[4])) & B[4];

	x = (x | (x << S[3])) & B[3];
	y = (y | (y << S[3])) & B[3];

	x = (x | (x << S[2])) & B[2];
	y = (y | (y << S[2])) & B[2];

	x = (x | (x << S[1])) & B[1];
	y = (y | (y << S[1])) & B[1];

	x = (x | (x << S[0])) & B[0];
	y = (y | (y << S[0])) & B[0];

	return x | (y << 1);
}

void geohashGetCoordRange(GeoHashRange *long_range, GeoHashRange *lat_range)
{
	/* These are constraints from EPSG:900913 / EPSG:3785 / OSGEO:41001 */
	/* We can't geocode at the north/south pole. */
	long_range->max = GEO_LONG_MAX;
	long_range->min = GEO_LONG_MIN;
	lat_range->max = GEO_LAT_MAX;
	lat_range->min = GEO_LAT_MIN;
}

int geohashEncode(const GeoHashRange *long_range, const GeoHashRange *lat_range,
				  double longitude, double latitude, uint8_t step,
				  GeoHashBits *hash)
{
	/* Check basic arguments sanity. */
	if (hash == NULL || step > 32 || step == 0 ||
		RANGEPISZERO(lat_range) || RANGEPISZERO(long_range))
		return 0;

	/* Return an error when trying to index outside the supported
     * constraints. */
	if (longitude > GEO_LONG_MAX || longitude < GEO_LONG_MIN ||
		latitude > GEO_LAT_MAX || latitude < GEO_LAT_MIN)
		return 0;

	hash->bits = 0;
	hash->step = step;

	if (latitude < lat_range->min || latitude > lat_range->max ||
		longitude < long_range->min || longitude > long_range->max)
	{
		return 0;
	}

	double lat_offset =
		(latitude - lat_range->min) / (lat_range->max - lat_range->min);
	double long_offset =
		(longitude - long_range->min) / (long_range->max - long_range->min);

	/* convert to fixed point based on the step size */
	lat_offset *= (1ULL << step);
	long_offset *= (1ULL << step);
	hash->bits = interleave64(lat_offset, long_offset);
	return 1;
}

int geohashEncodeType(double longitude, double latitude, uint8_t step, GeoHashBits *hash)
{
	GeoHashRange r[2] = {{0}};
	geohashGetCoordRange(&r[0], &r[1]);
	return geohashEncode(&r[0], &r[1], longitude, latitude, step, hash);
}

int geohashEncodeWGS84(double longitude, double latitude, uint8_t step,
					   GeoHashBits *hash)
{
	return geohashEncodeType(longitude, latitude, step, hash);
}

GeoHashFix52Bits geohashAlign52Bits(const GeoHashBits hash)
{
	uint64_t bits = hash.bits;
	bits <<= (52 - hash.step * 2);
	return bits;
}

int main()
{
	double xy[2];
	xy[0] = 30.26000;
	xy[1] = 120.18999;

	GeoHashBits hash;
	geohashEncodeWGS84(xy[0], xy[1], GEO_STEP_MAX, &hash);
	GeoHashFix52Bits bits = geohashAlign52Bits(hash);
	long long value = sdsfromlonglong(bits);
	printf("value: %lld\n", value);
	return 0;
}