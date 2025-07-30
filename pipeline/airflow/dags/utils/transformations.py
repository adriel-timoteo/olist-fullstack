import uuid
import pandas

def normalize_uuid(val):
    """Convert a value to a UUID object, if possible."""
    try:
        return uuid.UUID(str(val))
    except (ValueError, TypeError):
        return None

def clean_nulls(value):
    if pandas.isna(value):  # Handles NaN, NaT, None
        return None
    return value