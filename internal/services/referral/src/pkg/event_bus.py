"""Event bus implementation."""

from collections import defaultdict
from collections.abc import Callable

class EventBus:
    """Event bus implementation."""

    def __init__(self):
        """Initialize event bus."""
        self._listeners = defaultdict(list)

    def subscribe(self, event_type: str, listener: Callable):
        """Subscribe event."""
        self._listeners[event_type].append(listener)

    def unsubscribe(self, event_type: str, listener: Callable):
        """Unsubscribe event."""
        self._listeners[event_type].remove(listener)

    def publish(self, event_type: str, *args, **kwargs):
        """Publish event."""
        for listener in self._listeners[event_type]:
            listener(*args, **kwargs)
