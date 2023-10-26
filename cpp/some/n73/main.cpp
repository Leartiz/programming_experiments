#include <stdexcept>

#include <sstream>
#include <iostream>

using namespace std;

struct Date final
{
private:
    struct IoStateSaver final
    {
        explicit IoStateSaver(istream& in)
            : old{ in.exceptions() }, in{ in }
        {}
        ~IoStateSaver() {
            in.exceptions(old);
            cout << "~\n";
        }

        ios::iostate old;
        istream& in;
    };

public:
    int dd, mm, yy;
    explicit Date(istream& in) {
        if (in.fail()) throw invalid_argument("'in' is not good");
        IoStateSaver _(in);
        in.exceptions(ios::goodbit); // default iostate.

        char dots[]{ 0, 0 };
        in
            >> dd >> dots[0]
            >> mm >> dots[1]
            >> yy;
        if (in.fail()) throw runtime_error("'in' is failed");
        static_cast<void>(dots);

        // some checks...
    }

    void print(ostream& out) {
        out
            << dd << "."
            << mm << "."
            << yy;
    }
};

int main()
{
    auto sin = istringstream{ "12.12.2012" };
    auto flags = sin.exceptions();
    cout << flags << endl;

    try {
        Date date{ sin }; // temporary object not allowed!
        date.print(cout);
        cout << endl;
    }
    catch (exception& err) {
        cerr << "err: " << err.what() << '\n';
    }

    return 0;
}
