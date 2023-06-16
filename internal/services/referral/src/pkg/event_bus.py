"""Event bus implementation."""

from collections import defaultdict
from collections.abc import Callable
from typing import Any

class EventBus:
    """Event bus implementation."""

    def __init__(self):
        """Initialize event bus."""
        self._listeners = defaultdict(list)

    def subscribe(self, event_type: str, listener: Callable[..., Any]) -> None:
        """Subscribe event."""
        self._listeners[event_type].append(listener)

    def unsubscribe(self, event_type: str, listener: Callable[..., Any]) -> None:
        """Unsubscribe event."""
        if listener in self._listeners[event_type]:
            self._listeners[event_type].remove(listener)

    def publish(self, event_type: str, *args, **kwargs) -> list[Any]:
        """Publish event."""
        results = []
        for listener in self._listeners[event_type]:
            try:
                results.append(listener(*args, **kwargs))
            except Exception as e:
                print(f"Exception occurred while publishing event: {e}")
        return results
