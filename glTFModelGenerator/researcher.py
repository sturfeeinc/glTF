import json
from os import listdir
from os.path import isfile, join

schemasDir = "../shemas/"

onlyfiles = [f for f in listdir(schemasDir) if isfile(join(schemasDir, f))]

properties = dict()

for fileName in onlyfiles:
    schema = open(schemasDir + fileName)
    data = json.loads(schema.read())

    # that loop find all fields for schema level
    # for name in data:
    #     if name == "dependencies":
    #         print data["title"]
    #     properties[name] = 1

    # that loop find all fields for property level
    # if "properties" in data:
    #     for p in data['properties']:
    #         for c in data['properties'][p]:
    #             properties[c] = 1

    if "properties" in data:
        for p in data['properties']:
            for c in data['properties'][p]:
                if c == "type":
                    if "required" in data['properties'][p]:
                        print data['properties'][p][c]
                        print p


                    # try:
                    #     if data['properties'][p][c] == "number":
                    #         print p
                    #     properties[data['properties'][p][c]] = 1
                    # except:
                    #     print data['properties'][p][c]


for name in properties:
    print name
