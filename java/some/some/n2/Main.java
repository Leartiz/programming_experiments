package n2;

import java.lang.reflect.InvocationTargetException;

public class Main {
    public static void main(String[] args) {
        Report report = new Report("printHTML");

        try {
            report.print();

        } catch (InvocationTargetException e) {
            System.out.print("InvocationTargetException");
        } catch (IllegalAccessException e) {
            System.out.print("IllegalAccessException");
        } catch (NoSuchMethodException e) {
            System.out.print("NoSuchMethodException");
        }
    }
}
