package basic;

import edu.princeton.cs.algs4.In;
import org.junit.jupiter.api.Test;

class StackTest {

    @Test
    public void TestStack() {
        Stack<Integer> s = new Stack<Integer>();
        s.push(1);
        s.push(2);

        System.out.println(s.pop());
        System.out.println(s.pop());
    }

    @Test
    public void TestStackString() {
        Stack<String> s1 = new Stack<String>();
        s1.push("one");
        System.out.println(s1.size());
        s1.push("two");
        System.out.println(s1.size());
        s1.push("three");
        System.out.println(s1.size());

        System.out.println(s1.pop());
        System.out.println(s1.pop());
    }
}