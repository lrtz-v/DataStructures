package basic;

/**
 * 后进先出
 */
public class Stack<Item> {

    private Item[] a = (Item[]) new Object[1];
    private int N = 0;
    private int offset = 0;

    public boolean isEmpty() { return N == 0; }

    public int size() { return N; }

    private void resize(int cap) {
        Item[] temp = (Item[]) new Object[cap];
        for (int i = 0; i < N; i++) {
            temp[i] = a[i];
        }
        a = temp;
        offset = 0;
    }

    public void push (Item item) {
        if (N == a.length) { resize(a.length*2); }
        a[N++] = item;
    }

    public Item pop() {
        Item item = a[--N];
        a[N] = null;
        if (N > 0 && N == a.length/4) { resize(a.length/2); }
        return item;
    }
}
