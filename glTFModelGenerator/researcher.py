import json
from os import listdir
from os.path import isfile, join
onlyfiles = [f for f in listdir("shemas") if isfile(join("shemas", f))]


properties = dict()

for fileName in onlyfiles:
    schema = open("shemas/" + fileName)
    data = json.loads(schema.read())
    for name in data:
        if name == "dependencies":
            print data["title"]

        properties[name] = 1

for name in properties:
    print name