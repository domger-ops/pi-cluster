>make hello world webapp
- to run app in venv:
    sudo -E /path/to/venv/bin/activate
- Flask app Running locally on http://192.168.254.164:5000

>build and push docker image
- created dockerfile

>deploy to cluster in new namespace
- ns = webapp 
- fixed crashloopbackoff with requirements.txt

>create phonebook webapp
- create db w/ postgresql. 
- to login> psql -U postgres
- flask commands...
    source venv/bin/activate
    export FLASK_APP=phonebook.py
    pip install flask-login
    pip install psycopg2-binary
~ docker compose for multiple resources/containers ie.database/etc


>deploy phonebook app to the cluster instead of helloworld webapp


>secure and link cloudk.ing to host ip
steps to secure: