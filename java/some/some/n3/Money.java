package n3;

public class Money {
    private double value = 0;

    private void setValue(double value) {
        this.value = value;
    }

    public double getValue() {
        return value;
    }

    public Money(double value) {
        this.value = value;
    }

    public Money sum(Money money) {
        Money newMoney = Money.zero();
        newMoney.setValue(
                money.getValue() +
                        getValue());
        return newMoney;
    }

    public static Money zero() {
        return new Money(0);
    }
}