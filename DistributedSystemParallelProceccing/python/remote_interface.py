import abc

class RemoteInterface(abc.ABC):
    @abc.abstractmethod
    def say_hello(self) -> str:
       pass
    @abc.abstractmethod
    def add(self,a,b) -> int:
       pass