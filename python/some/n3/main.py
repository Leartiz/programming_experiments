# 21 глава

class WasRun:
    def __init__(self, name):
        self.name = name
        self.log = ""

    def set_up(self):
        self.log = "set_up "

    def test_method(self):
        self.log = self.log + "test_method "

    def tear_down(self):
        self.log = self.log + "tear_down "

    # *** ?

    def run(self):
        result = TestResult()
        result.test_started()

        self.set_up()
        method = getattr(self, self.name)
        method()
        self.tear_down()

        return result


class TestCase:
    def __init__(self, name):
        self.name = name
        self.test = None

    def set_up(self):
        pass

    def run(self):
        self.set_up()
        exec("self." + self.name + "()")
        self.tear_down()

    def tear_down(self):
        pass


class TestCaseTest(TestCase):
    def set_up(self):
        pass

    def test_template_method(self):
        test = WasRun("test_method")
        test.run()
        assert test.log == "set_up test_method tear_down "

    def test_result(self):
        test = WasRun("test_method")
        test.run()
        assert test.log == "set_up test_method tear_down "

    def tear_down(self):
        pass


class TestResult:
    def __init__(self):
        self.run_count = 1

    def test_started(self):
        self.run_count = self.run_count + 1

    def summary(self):
        return "%d run, 0failed" % self.run_count


if __name__ == '__main__':
    TestCaseTest("test_template_method").run()