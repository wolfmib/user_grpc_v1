## Find method in golang code

---

- cursor, err := collection.Find(ctx, bson.M{"firstname": "hi"})

---

var \_key string = "firstname"
var \_value string = my_user.FirstName

---

cursor, err := collection.Find(ctx, bson.M{\_key: \_value})

---
