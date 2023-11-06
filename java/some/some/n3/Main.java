package n3;

// Composite pattern.

public class Main {
    public static void main(String[] args) {
        Account account = new Account();
        account.addHolding(new Transaction(new Money(100)));
        account.addHolding(new Transaction(new Money(111)));
        account.addHolding(new Transaction(new Money(7)));
        account.addHolding(generateFakeAccount());
        //...

        System.out.print(account.balance().getValue());
    }

    public static Account generateFakeAccount() {
        Account account = new Account();
        account.addHolding(new Transaction(new Money(1)));
        account.addHolding(new Transaction(new Money(1000)));
        account.addHolding(new Transaction(new Money(42)));
        return account;
    }
}
