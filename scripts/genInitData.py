import faker
from random import choice, randint

filesLocation = "build/package/postgres/"

userFile = filesLocation + "user.csv"
brandFile = filesLocation + "brand.csv"
itemFile = filesLocation + "item.csv"
orderingFile = filesLocation + "ordering.csv"
orderItemsFile = filesLocation + "orderItems.csv"

USER_ROWS = 3000
BRAND_ROWS = 1000
ITEM_ROWS = 5000
ORDERING_ROWS = 10000
ORDER_ITEMS_ROWS = 50000

myFaker = faker.Faker("en_US")
myRuFaker = faker.Faker("ru_RU")

def generateUser():
    sex = ["male", "female"]
    file = open(userFile, "w", encoding="utf-8")
    role = ["guest", "user", "admin"]

    for i in range(USER_ROWS):
        curSex = choice(sex)

        if "female" == curSex:
            firstName = myFaker.first_name_female()
        else:
            firstName = myFaker.first_name_male()

        line = "{};{};{};{};{};{}\n".format(i + 1, myFaker.unique.user_name(), myFaker.password(), firstName, curSex, choice(role))

        file.write(line)

    file.close()

def generateBrand():
    file = open(brandFile, "w", encoding="utf-8")

    for i in range(BRAND_ROWS):
        line = "{};{};{};{};{}\n".format(i + 1, myFaker.company(), myFaker.year(), myFaker.iana_id(), myFaker.company())

        file.write(line)

    file.close()

def generateItem():
    file = open(itemFile, "w", encoding="utf-8")
    category = ["ботинки", "кроссовки", "майка", "футболка", "куртка", "штаны", "шорты", "ремень", "шляпа"]
    size = ["XS", "S", "M", "L", "XL", "XXL"]
    sex = ["male", "female"]
    boolean = ["true", "false"]

    for i in range(ITEM_ROWS):
        line = "{};{};{};{};{};{};{};{}\n".format(i + 1, choice(category), choice(size), randint(100, 20000), 
                                            choice(sex), myFaker.iana_id(), randint(1, BRAND_ROWS), choice(boolean))

        file.write(line)

    file.close()

def generateOrdering():
    file = open(orderingFile, "w", encoding="utf-8")
    status = ["оформлен", "доставлен", "отменен"]

    for i in range(ORDERING_ROWS - USER_ROWS):
        line = "{};{};{};{};{}\n".format(i + 1, myFaker.date_this_century(), randint(1, USER_ROWS), randint(100, 15000), 
                                            choice(status))

        file.write(line)

    for i in range(ORDERING_ROWS - USER_ROWS, ORDERING_ROWS):
        line = "{};{};{};{};{}\n".format(i + 1, "null", i - (ORDERING_ROWS - USER_ROWS) + 1, "null", 
                                            "корзина")

        file.write(line)        

    file.close()

def generateOrderItems():
    file = open(orderItemsFile, "w", encoding="utf-8")
    status = ["корзина", "оформлен", "доставлен", "отменен"]

    for i in range(ORDER_ITEMS_ROWS):
        line = "{};{};{};{}\n".format(i + 1, randint(1, ORDERING_ROWS), randint(1, ITEM_ROWS), randint(1, 3))

        file.write(line)

    file.close()


generateUser()
generateBrand()
generateItem()
generateOrdering()
generateOrderItems()
