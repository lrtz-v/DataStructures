package basic;
import org.junit.jupiter.api.Test;
import edu.princeton.cs.algs4.StdRandom;


class CounterTest {

    @Test
    public void TestCounter() {
        Counter test = new Counter("test");
        test.increment();
        System.out.println(test.tally());
        test.increment();
        System.out.println(test.tally());
        System.out.println(test);
    }

    @Test
    public void TestCounter2() {
        int T = 10;
        Counter count1 = new Counter("one");
        Counter count2 = new Counter("two");
        for (int t = 0; t < T; t++) {
            if (StdRandom.bernoulli(0.5)) {
                count1.increment();
            } else {
                count2.increment();
            }
        }

        System.out.println(count1);
        System.out.println(count2);

        int d = count1.tally() - count2.tally();
        System.out.println("delta: "+ Math.abs(d));
    }

}