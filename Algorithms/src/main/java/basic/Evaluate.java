package basic;

public class Evaluate {

    public static int evaluate(String evaluateStr) {
        basic.Stack<String> ops = new basic.Stack<String>();
        basic.Stack<Integer> vals = new basic.Stack<Integer>();

        for (int i = 0; i < evaluateStr.length(); i++) {
            String s = String.valueOf(evaluateStr.charAt(i));
            if (s.equals("(")) {}
            else if (s.equals("+")) {ops.push(s);}
            else if (s.equals("-")) {ops.push(s);}
            else if (s.equals("*")) {ops.push(s);}
            else if (s.equals("/")) {ops.push(s);}
            else if (s.equals(")")) {
                String op = ops.pop();
                int v = vals.pop();

                if (op.equals("+")) {v = vals.pop() + v;}
                else if (op.equals("-")) {v = vals.pop() - v;}
                else if (op.equals("*")) {v = vals.pop() * v;}
                else if (op.equals("/")) {v = vals.pop() / v;}
                vals.push(v);
            }
            else {
                vals.push(Integer.parseInt(s));
            }
        }
        return vals.pop();
    }
}
