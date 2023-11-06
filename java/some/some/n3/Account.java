package n3;

import java.util.Arrays;

public class Account implements Holding{
    private Holding holdings[] = new Holding[0];

    public void addHolding(Holding holding) {
        holdings = Arrays.copyOf(holdings, holdings.length + 1);
        holdings[holdings.length - 1] = holding;
    }

    @Override
    public Money balance() {
        Money sum = Money.zero();
        for (int i = 0; i < holdings.length; ++i)
            sum = sum.sum(holdings[i].balance());
        return sum;
    }
}
