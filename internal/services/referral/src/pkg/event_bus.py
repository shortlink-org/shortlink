from collections import defaultdict
from typing import Callable

class EventBus:
    def __init__(self):
        self._listeners = defaultdict(list)

    def subscribe(self, event_type: str, listener: Callable):
        self._listeners[event_type].append(listener)

    def unsubscribe(self, event_type: str, listener: Callable):
        self._listeners[event_type].remove(listener)

    def publish(self, event_type: str, *args, **kwargs):
        for listener in self._listeners[event_type]:
            listener(*args, **kwargs)
