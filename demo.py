#!coding: utf-8

class Foo:
    def __init__(self, aInt, bStr):
        self.a = aInt
        self.b = bStr
        print("__init__ is called: %d, %s\n" % (aInt, bStr))

    def invoke(self, reqInt):
        print("invoke is called: reqInt: %s\n" % reqInt)
        return self.a + reqInt
