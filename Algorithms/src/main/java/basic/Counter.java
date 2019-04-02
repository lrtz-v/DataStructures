package basic;

public class Counter {

    private final String name;
    private int count;

    public Counter(String counterName) {
        name = counterName;
        count = 0;
    }

    public void increment() {
        count++;
    }

    public int tally() {
        return count;
    }

    @Override
    public String toString() {
        return "Counter: " + name + ", Count: " + count;
    }
}
