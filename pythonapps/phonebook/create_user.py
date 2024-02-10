import os
from phonebook import db, User
from secrets import USER_USERNAME, USER_PASSWORD

def create_user(username, password):
    # Create a new user instance and add it to the database
    new_user = User(username=username, password=password)
    
    # Add the user to the database session
    db.session.add(new_user)
    
    # Commit the transaction
    db.session.commit()

if __name__ == "__main__":
    # Get username and password from secrets file
    username = USER_USERNAME
    password = USER_PASSWORD
    
    # Call the function to create the user
    create_user(username, password)
    print("User created successfully!")
