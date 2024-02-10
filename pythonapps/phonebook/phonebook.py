from flask import Flask, request, jsonify, render_template, redirect, url_for
from flask_sqlalchemy import SQLAlchemy
from flask_login import UserMixin, login_user, LoginManager, login_required, current_user, logout_user
from dotenv import load_dotenv
import os

app = Flask(__name__)
load_dotenv()

# Configure the database
app.config['SQLALCHEMY_DATABASE_URI'] = f"postgresql://{os.getenv('DB_USERNAME')}:{os.getenv('DB_PASSWORD')}@localhost/phonebook"
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False
app.config['SECRET_KEY'] = 'your_secret_key'  # Change this to a random secret key
db = SQLAlchemy(app)

# Import models
import models

# Flask-Login configuration
login_manager = LoginManager()
login_manager.init_app(app)

@login_manager.user_loader
def load_user(user_id):
    return models.User.query.get(int(user_id))

# Define User model
class User(UserMixin, db.Model):
    id = db.Column(db.Integer, primary_key=True)
    username = db.Column(db.String(100), unique=True, nullable=False)
    password = db.Column(db.String(100), nullable=False)

# Define a route for the root endpoint
@app.route('/')
def index():
    return render_template('index.html')

# Route for adding a contact
@app.route('/add_contact', methods=['POST'])
def add_contact():
    data = request.get_json()
    new_contact = models.Contact(name=data['name'], phone_number=data['phone_number'], email=data['email'])
    db.session.add(new_contact)
    db.session.commit()
    return jsonify({'message': 'Contact added successfully'}), 201

# Route for listing contacts (accessible only after authentication)
@app.route('/list_contacts')
@login_required
def list_contacts():
    # Implement filtering logic
    name_filter = request.args.get('name')
    if name_filter:
        contacts = models.Contact.query.filter(models.Contact.name.ilike(f"%{name_filter}%")).all()
    else:
        contacts = models.Contact.query.all()
    contact_list = [{'name': contact.name, 'phone_number': contact.phone_number, 'email': contact.email} for contact in contacts]
    return render_template('list_contacts.html', contacts=contact_list)

# Login route
@app.route('/login', methods=['GET', 'POST'])
def login():
    if request.method == 'POST':
        username = request.form['username']
        password = request.form['password']
        user = User.query.filter_by(username=username, password=password).first()
        if user:
            login_user(user)
            return redirect(url_for('list_contacts'))
        else:
            return 'Invalid username or password'
    else:
        return render_template('login.html')

# Logout route
@app.route('/logout')
@login_required
def logout():
    logout_user()
    return redirect(url_for('index'))

if __name__ == "__main__":
    # Run the Flask app on port 5000
    app.run(port=5000)