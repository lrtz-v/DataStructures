#include "stdint.h"

#define GEO_LAT_MIN -85.05112878
#define GEO_LAT_MAX 85.05112878
#define GEO_LONG_MIN -180
#define GEO_LONG_MAX 180

#define GEO_STEP_MAX 26 /* 26*2 = 52 bits. */

typedef uint64_t GeoHashFix52Bits;

#define HASHISZERO(r) (!(r).bits && !(r).step)
#define RANGEISZERO(r) (!(r).max && !(r).min)
#define RANGEPISZERO(r) (r == NULL || RANGEISZERO(*r))

typedef struct
{
    uint64_t bits;
    uint8_t step;
} GeoHashBits;

typedef struct
{
    double min;
    double max;
} GeoHashRange;

typedef struct
{
    GeoHashBits hash;
    GeoHashRange longitude;
    GeoHashRange latitude;
} GeoHashArea;

typedef struct
{
    GeoHashBits north;
    GeoHashBits east;
    GeoHashBits west;
    GeoHashBits south;
    GeoHashBits north_east;
    GeoHashBits south_east;
    GeoHashBits north_west;
    GeoHashBits south_west;
} GeoHashNeighbors;