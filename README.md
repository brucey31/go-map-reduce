## Calculation to be attempted
Given a huge JSON of basket objects, use map-reduce to create a total list of different items
e.g.

```json
[
   {
      "basket owner":"422dbd27-61ef-4dd2-9404-345b51679da8",
      "contents":[
         "apple",
         "apple",
         "banana",
         "cucumber",
         "corgette",
         "turnip"
      ]
   },
   {
      "basket owner":"a8bd018e-5a69-4741-bcb5-1aedd314453f",
      "contents":[
         "cabbage",
         "pepper",
         "pepper",
         "mushrooms",
         "kale",
         "spinach"
      ]
   }
]
```

## Compile protos
```bash
protoc --go_out=plugins=grpc:. src/calculator/v1/service.proto
```

### Issues & Solution links
- How to import from other packages [link](https://www.digitalocean.com/community/tutorials/how-to-use-go-modules)