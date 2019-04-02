package basic;

import org.junit.jupiter.api.Test;

class TrainingATest {

    @Test
    void lnOfFactorial() {
        System.out.print(TrainingA.lnOfFactorial(10));
    }

    @Test
    void statisticsFile() {
        TrainingA.statisticsFile(System.getProperty("user.dir")+"/src/main/java/basic/trainingA.text");
    }

    @Test
    void euclid() {
        System.out.printf("num1: %d\tnum2:%d\tres:%d\n", 105, 24, TrainingA.euclid(105, 24));
        System.out.printf("num1: %d\tnum2:%d\tres:%d\n", 24, 105, TrainingA.euclid(24, 105));
    }

    @Test
    void euclid2() {
        System.out.printf("num1: %d\tnum2:%d\tres:%d\n", 1111111, 1234567, TrainingA.euclid(1111111, 1234567));
    }

    @Test
    void createMatrix() {
        boolean[][] a = TrainingA.createMatrix(4);
        for(boolean[] x : a) System.out.printf("%b\t%b\t%b\t%b\n", x[0], x[1], x[2], x[3]);
    }

    @Test
    void drawCycle() {
        TrainingA.drawCycle(0.5, 0.5, 0.2);
    }
}