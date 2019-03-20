package basic;
import java.io.*;
import edu.princeton.cs.algs4.StdDraw;

/**
 * TrainingA
 * @author tao
 */
public class TrainingA {

    /**
     * ln(N!)
     **/
    public static double lnOfFactorial(int n) {

        if (n == 0) {return 1;}
        return Math.log1p(n)*Math.log1p(n-1);
    }

    /**
     *
     * @param filename
     *
     * file:
     *  Bob 10 15
     *  Alice 7 10
     *  Linda 19 3
     */
    public static void statisticsFile(String filename) {
        try {
            FileReader reader = new FileReader(filename);
            BufferedReader br = new BufferedReader(reader);

            String line;
            while ((line = br.readLine()) != null) {
                String[] vals = line.split(" ");
                String name = vals[0];
                int score1 = Integer.parseInt(vals[1]);
                int score2 = Integer.parseInt(vals[2]);
                System.out.printf("Name: %s\tScore1: %d\tScore2: %d\tRange: %.3f\n", name, score1, score2, 1.0*score1/score2);
            }

        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    /**
     * 最大公约数
     * (p, q) = (q, p % q)
     */
    public static int euclid(int p, int q) {
        if (q == 0) {return p;}
        return euclid(q, p % q);
    }

    /**
     * 创建bool NxN a[][], ij互质时, 值为true, 否则为false
     *
     */
    public static boolean[][] createMatrix(int n) {
        boolean[][] a = new boolean[n][n];
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                a[i][j] = euclid(i, j) == 1;
            }
        }
        return a;
    }

    /**
     * 画一个圆
     */
    public static void drawCycle(double x, double y, double r) {
        StdDraw.circle(x, y, r);
        StdDraw.show();
    }
}
