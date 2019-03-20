package basic;

import org.junit.Assert;
import org.junit.jupiter.api.Test;

class MatrixTest {

    @Test
    void dot() {
        int[] a = {1, 2, 3};
        int res = Matrix.dot(a, a);
        int expected = 14;
        Assert.assertEquals(expected, res);
    }

    @Test
    void transpose() {
        int[][] a = {{1, 0, 3}, {2, 1, 0}};
        Matrix.displayMatrix(Matrix.transpose(a));
    }

    @Test
    void multNum() {
        int[][] a = {{1, 0, 3}, {2, 1, 0}};
        Matrix.displayMatrix(Matrix.multNum(a, 2));
    }

    @Test
    void mult() {
        int[][] a = {{1, 0, 2}, {-1, 3, 1}};
        int[][] b = {{3, 1}, {2, 1}, {1, 0}};
        Matrix.displayMatrix(Matrix.mult(a, b));
    }

    @Test
    void add() {
        int[][] a = {{1, 3, 1}, {1, 0, 0}};
        int[][] b = {{0, 0, 5}, {7, 5, 0}};
        Matrix.displayMatrix(Matrix.add(a, b));
    }

}