package basic;
import java.util.ArrayList;
import java.util.List;

import org.junit.Assert;
import org.junit.jupiter.api.Test;

class TwoSumTest {

    @Test
    void twoSum() {
        List<Integer> args = new ArrayList<Integer>(); // [1, 8, 3, 3, 2]
        int target = 3;
        args.add(1);
        args.add(8);
        args.add(3);
        args.add(3);
        args.add(2);

        List<Integer> res = TwoSum.twoSum(args, target);

        List<Integer> expect_res = new ArrayList<Integer>();
        expect_res.add(0);
        expect_res.add(4);

        Assert.assertArrayEquals(String.format("[*] Expect: %s, Get: %s", expect_res, res),
                expect_res.toArray(),
                res.toArray());
    }

    @Test
    void twoSum2() {
        List<Integer> args = new ArrayList<Integer>(); // [1, 1, 2]
        int target = 3;
        args.add(1);
        args.add(1);
        args.add(2);

        List<Integer> res = TwoSum.twoSum(args, target);

        List<Integer> expect_res = new ArrayList<Integer>();
        expect_res.add(0);
        expect_res.add(2);

        Assert.assertArrayEquals(String.format("[*] Expect: %s, Get: %s", expect_res, res),
                expect_res.toArray(),
                res.toArray());
    }
}