#!/usr/bin/python3
import requests
import sys

class Menu:

    username = ""
    password = ""
    server = ""
    def create_dropdown_menu(self, text, choices):
        print(text)
        for x in range(len(choices)):
            print("{}) {}".format(x+1, choices[x]))
        return input("What is your choice?: ")

    def question_menu(self, text):
        return input(text + ": ")

    def login_menu(self):
        while True:
            self.username = menu.question_menu("what is your username?")
            self.password = menu.question_menu("what is your password?")
            self.server = menu.question_menu("whats the server address?")
            if requests.post(self.server+"/verifyUser", auth=(self.username, self.password)).status_code == 200:
                return

    def main_menu(self):
        user_input = self.create_dropdown_menu("hello! ", [
            "Create Post",
            "See Posts",
            "Edit Post",
            "Quit"
            ])
        if user_input == "1":
            self.create_post_menu()

        elif user_input == "2":
            json_data = requests.get(self.server+"/getAllPosts", auth=(self.username, self.password)).json()["posts"]
            for x in json_data:
                print("\nmood: {}  genre: {}  time: {}\n{}".format(x["mood"], x["genre"], x["timestamp"], x["post"]))

        elif user_input == "4":
            sys.exit()
        else:
            print("enter something from the menu please!")

    def create_post_menu(self):
        mood = self.question_menu("how are you feeling?")
        genre = self.question_menu("What is this about?")
        post = input("please type post below:\n")
        requests.post(self.server + "/addPost", 
                json={"username":self.username, "mood": mood, "genre": genre, "post": post},
                auth=(self.username, self.password)
                )

if __name__ == "__main__":
    menu = Menu()
    menu.login_menu()
    while(True):
        menu.main_menu()
