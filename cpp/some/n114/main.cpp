#include <cstdio>
#include <string>

using namespace std;

struct UnitComponent {
    string name;
    int number;

    UnitComponent(string nm, int num) : name(nm), number(num) {}
    virtual bool operator==(const UnitComponent& a) const {
        printf("bool operator==(const UnitComponent&)\n");
        return name == a.name && number == a.number;
    }
    // ...
};

struct UnitHand : public UnitComponent {
    bool left_hand;

    UnitHand(string nm, int num, char ch) : UnitComponent(nm, num), left_hand(ch) {}
    virtual bool operator==(const UnitHand& a) const {
        printf("bool operator==(const UnitHand&)\n");
        return UnitComponent::operator==(a) && left_hand == a.left_hand;
    }
    // ...
};

int main() {
    UnitComponent comp{ "test", 1 };
    UnitHand hand{ "test", 1, true };
    printf("comp == hand (%d)\n", comp == hand); // compares name and number, ignores unit's left_hand

    UnitHand leg{ "test", 1, false };
    printf("hand == leg (%d)\n", hand == leg); // compares name, number, and left_hand

    UnitComponent &hand_ref = hand;
    UnitComponent &leg_ref = leg;
    printf("hand_ref == leg_ref (%d)\n",
           hand_ref == leg_ref);  // ???

    return 0;
}
