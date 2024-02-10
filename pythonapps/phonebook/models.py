from sqlalchemy import CheckConstraint
from phonebook import db
import re

class Contact(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(100), nullable=False)
    phone_number = db.Column(db.String(10), nullable=False, unique=True)
    email = db.Column(db.String(100), nullable=False, unique=True)

    def __repr__(self):
        return f"Contact('{self.name}', '{self.phone_number}', '{self.email}')"

# Check constraint for phone number format (10 digits only)
CheckConstraint('LENGTH(phone_number) = 10 AND phone_number SIMILAR TO \'[0-9]{10}\'',
                name='valid_phone_number'),

# Check constraint for email format (simple email format validation)
CheckConstraint('email ~* \'^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$\'',
                name='valid_email')
