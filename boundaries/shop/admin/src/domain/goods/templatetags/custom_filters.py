"""This module contains custom template filters for the goods domain.

These filters are used to extend the template language with custom functionality specific to the goods domain.
"""

from django import template

register = template.Library()


@register.filter(name="length_is")
def length_is(value, length):
    """Check if the length of the given value is equal to the specified length.

    Args:
        value: The value whose length is to be checked.
        length: The length to compare against.

    Returns:
        bool: True if the length of the value is equal to the specified length, False otherwise.
    """
    return len(value) == length
