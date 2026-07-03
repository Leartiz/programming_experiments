#include <gtest/gtest.h>

// TEST_P
// ----------------------------------------------

class LeapYearCalendar final {
public:
    bool isLeap(const int value)
    {
        if (value % 4 != 0) {
            return false;
        }

        if (value % 100 != 0) {
            return true;
        }

        if (value % 400 != 0) {
            return false;
        }
        
        return true;
    }
};

class LeapYearMultipleParametersTests : public ::testing::TestWithParam<std::tuple<int, bool>> {
protected:
    LeapYearMultipleParametersTests()
    {
        static int count = 0;
        std::cerr << "Ctor LeapYearMultipleParametersTests. "
                  << "Count: " << count++ << std::endl;
    }

protected:
    LeapYearCalendar leapYearCalendar;
};
    
TEST_P(LeapYearMultipleParametersTests, ChecksIfLeapYear) {
    bool expected = std::get<1>(GetParam());
    int year = std::get<0>(GetParam());
    ASSERT_EQ(expected, leapYearCalendar.isLeap(year));
}

INSTANTIATE_TEST_CASE_P(
    LeapYearTests,
    LeapYearMultipleParametersTests,
    ::testing::Values(
        std::make_tuple(7, false),
        std::make_tuple(2001, false),
        std::make_tuple(1996, true),
        std::make_tuple(1700, false),
        std::make_tuple(1600, true)
    )
);

// TEST_F
// ----------------------------------------------

class LeapYearFixture : public testing::Test {
protected:
    LeapYearFixture()
    {
        static int count = 0;
        std::cerr << "Ctor LeapYearFixture. "
                  << "Count: " << count++ << std::endl;
    }

protected:
    LeapYearCalendar leapYearCalendar;
};
    
TEST_F(LeapYearFixture, test_7_false)
{
    bool expected = false;
    int year = 7;
    ASSERT_EQ(expected, leapYearCalendar.isLeap(year));
}

TEST_F(LeapYearFixture, test_2001_false)
{
    bool expected = false;
    int year = 2001;
    ASSERT_EQ(expected, leapYearCalendar.isLeap(year));
}

TEST_F(LeapYearFixture, test_1996_true)
{
    bool expected = true;
    int year = 1996;
    ASSERT_EQ(expected, leapYearCalendar.isLeap(year));
}
