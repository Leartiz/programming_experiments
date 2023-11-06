package n1;

public class Report {
    String printMessage;

    public Report(String printMessage) {
        this.printMessage = printMessage;
    }

    // ---> Pluggable Selector
    public void print() {
        switch (printMessage) {
            case "printHTML":
                printHTML();
                break;
            case "printXML":
                printXML();
                break;
        }
    }
    void printHTML() {
        System.out.print("HTML");
    }
    void printXML() {
        System.out.print("XML");
    }
}