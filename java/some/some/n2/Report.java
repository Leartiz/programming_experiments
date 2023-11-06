package n2;

import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;

public class Report {
    String printMessage;

    public Report(String printMessage) {
        this.printMessage = printMessage;
    }

    public void print() throws InvocationTargetException, IllegalAccessException, NoSuchMethodException {
        Method runMethod = getClass().getDeclaredMethod(printMessage);
        runMethod.invoke(this);
    }
    void printHTML() {
        System.out.print("HTML");
    }
    void printXML() {
        System.out.print("XML");
    }
}