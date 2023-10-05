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
        self.set_up()
        method = getattr(self, self.name)
        method()
        self.tear_down()


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
        self.test = WasRun("test_method")

    def test_template_method(self):
        self.test.run()
        assert self.test.log == "set_up test_method tear_down "

    def tear_down(self):
        self.test = None


if __name__ == '__main__':
    TestCaseTest("test_template_method").run()