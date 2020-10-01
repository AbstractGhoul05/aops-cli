import mechanicalsoup
import argparse

browser = mechanicalsoup.StatefulBrowser()
browser.open("https://artofproblemsolving.com/?login=1")

browser.select_form('#login-form')
browser.get_current_form().print_summary()
