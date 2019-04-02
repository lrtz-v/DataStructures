package basic;

import org.junit.Assert;
import org.junit.jupiter.api.Test;

class BinarySearchTest {

    @Test
    void rank() {
        int[] a = {1, 2, 3, 4, 5, 6, 7, 8};
        Assert.assertEquals(3, BinarySearch.rank(a, 4));
        Assert.assertEquals(7, BinarySearch.rank(a, 8));
        Assert.assertEquals(-1, BinarySearch.rank(a, 9));
    }

    @Test
    void countSmaller() {
        int[] a = {1, 2, 3, 4, 5, 6, 7, 8};
        Assert.assertEquals(3, BinarySearch.countSmaller(a, 4));
        Assert.assertEquals(-1, BinarySearch.countSmaller(a, 9));
    }

    @Test
    void countGreater() {
        int[] a = {1, 2, 3, 4, 5, 6, 7, 8};
        Assert.assertEquals(3, BinarySearch.countGreater(a, 5));
        Assert.assertEquals(7, BinarySearch.countGreater(a, 1));
    }

    @Test
    void countEqual() {
        int[] a = {1, 2, 3, 4, 5, 6, 7, 8};
        Assert.assertEquals(1, BinarySearch.countEqual(a, 5));
        Assert.assertEquals(1, BinarySearch.countEqual(a, 6));

        int[] b = {1, 2, 3, 5, 5, 5, 7, 8};
        Assert.assertEquals(3, BinarySearch.countEqual(b, 5));
    }
}