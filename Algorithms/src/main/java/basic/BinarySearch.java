package basic;

/**
 * BinarySearch
 * @author tao
 */
public class BinarySearch {

    /**
     *
     * @param a: sorted list
     * @param val: value to find
     * @return index of val
     */
    public static int rank(int[] a, int val) {
        int lo = 0;
        int hi = a.length - 1;
        while (lo <= hi) {
            int mid = lo + (hi - lo)/2;
            if (val < a[mid]) {hi = mid - 1;}
            else if (val > a[mid]) {lo = mid + 1;}
            else {return mid;}
        }
        return -1;
    }

    /**
     *
     * @param a: sorted list
     * @param val: value to find
     * @return count of entries smaller than val
     */
    public static int countSmaller(int[] a, int val) {
        int index = rank(a, val);
        if (index < 0) {return -1;}
        if (index == 0) {return 0;}

        for (int i = index - 1; i >= 0; i--) {
            if (a[i] != val) {return i+1;}
        }
        return 0;
    }

    /**
     * {1, 2, 3, 4, 5, 6, 7, 8}
     * @param a: sorted list
     * @param val: value to find
     * @return count of entries smaller than val
     */
    public static int countGreater(int[] a, int val) {
        int index = rank(a, val);
        if (index < 0) {return -1;}
        if (index == 0) {return a.length-1;}

        for (int i = index + 1; i < a.length; i++) {
            if (a[i] != val) {return a.length-i;}
        }
        return 0;
    }

    /**
     * {1, 2, 3, 4, 5, 6, 7, 8}
     * @param a: sorted list
     * @param val: value to find
     * @return count of entries equal to val
     */
    public static int countEqual(int[] a, int val) {
        int smaller = countSmaller(a, val);
        int greater = countGreater(a, val);

        if (smaller >= 0 && greater >= 0) {
            return a.length - smaller - greater;
        }
        return 0;
    }
}
