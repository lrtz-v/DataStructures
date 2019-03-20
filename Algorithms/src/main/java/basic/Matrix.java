package basic;

public class Matrix {

    /**
     * 向量点乘
     * a = [a1, a2, a3, a4]
     * b = [b1, b2, b3, b4]
     * a*b = a1*b1 + a2*b2 + a3*b3 + a4*b4
     */
    public static int dot(int[] x, int[] y) {
        int res = 0;
        for (int i = 0; i < x.length; i++) {
            res += x[i]*y[i];
        }
        return res;
    }

    /**
     * 转制矩阵
     * [[1, 0, 3],   ===>  [[1, 2],
     *  [2, 1, 0]]   ===>   [0, 1],
     *                      [3, 0]]
     */
    public static int[][] transpose(int[][] a) {
        int[][] b = new int[a[0].length][a.length];
        for (int i = 0; i < a.length; i++) {
            for (int j = 0; j < a[0].length; j++) {
                b[j][i] = a[i][j];
            }
        }
        return b;
    }

    /**
     * 矩阵与数字之积
     */
    public static int[][] multNum(int[][] a, int n) {
        int[][] b = new int[a.length][a[0].length];
        for (int i = 0; i < a.length; i++) {
            for (int j = 0; j < a[0].length; j++) {
//                a[i][j] *= n;
                b[i][j] = a[i][j]*n;
            }
        }
        return b;
    }

    /**
     * 矩阵与矩阵之积
     * 两个矩阵的乘法仅当第一个矩阵的列数(column)和另一个矩阵的行数(row)相等时才能定义。
     * 如A是m x n矩阵和B是n x p矩阵，它们的乘积AB是一个m x p矩阵
     * a = {{1, 0, 2}, {-1, 3, 1}}
     * b = {{3, 1}, {2, 1}, {1, 0}}
     * res: [[5, 1], [4, 2]]
     */
    public static int[][] mult(int[][] a, int[][] b) {
        int row = a.length;
        int col = b[0].length;
        int[][] res = new int[row][col];
        for (int i = 0; i < row; i++) {
            for (int j = 0; j < col; j++) {
                for (int k = 0; k < b.length; k++) {
                    res[i][j] += a[i][k] * b[k][j];
                }
            }
        }

        return res;
    }

    /**
     * 矩阵与矩阵之和
     * a = {{1, 3, 1}, {1, 0, 0}}
     * b = {{0, 0, 5}, {7, 5, 0}}
     * res: [[1, 3, 6], [8, 5, 0]]
     */
    public static int[][] add(int[][] a, int[][] b) {
        int[][] res = new int[a.length][a[0].length];
        for (int i = 0; i < a.length; i++) {
            for (int j = 0; j < a[0].length; j++) {
                res[i][j] = a[i][j] + b[i][j];
            }
        }
        return res;
    }

    public static void displayMatrix(int[][] product) {
        System.out.println("Get: ");
        for(int[] row : product) {
            for (int column : row) {
                System.out.printf("%d\t", column);
            }
            System.out.println();
        }
    }

    public static void displayArray(int[] product) {
        System.out.println("Get: ");
        for(int x : product) {
            System.out.printf("%d\t", x);
        }
        System.out.println();
    }
}
