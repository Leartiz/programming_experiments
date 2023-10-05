class WasRun:
    def __init__(self, name):
        self.wasSetUp = None
        self.wasRun = None
        self.name = name

    def set_up(self):
        self.wasSetUp = 1

    def test_method(self):
        self.wasRun = 1

    # ***

    def run(self):
        method = getattr(self, self.name)
        method()


class TestCase:
    def __init__(self, name):
        self.name = name
        self.test = None

    def set_up(self):
        pass

    def run(self):
        self.set_up()
        method = getattr(self, self.name)
        method()


class TestCaseTest(TestCase):
    def set_up(self):
        self.test = WasRun("test_method")

    def test_set_up(self):
        self.test.set_up()
        assert self.test.wasSetUp

    def test_running(self):
        self.test.run()
        assert self.test.wasRun


if __name__ == '__main__':
    TestCaseTest("test_set_up").run()
    TestCaseTest("test_running").run()


