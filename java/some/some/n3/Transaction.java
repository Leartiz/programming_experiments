package n3;

public class Transaction implements Holding {
    private Money value;

    public Transaction(Money value) {
        this.value = value;
    }

    @Override
    public Money balance() {
        return value;
    }
}
