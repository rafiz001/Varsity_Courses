import abc

class RemoteInterface(abc.ABC):
    @abc.abstractmethod
    def say_hello(self) -> str:
       pass