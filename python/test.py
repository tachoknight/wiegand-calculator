#!/usr/bin/env python3

import wiegand


def main():
    print("Wiegand test")
    tag_to_convert = 2802899
    print(f"Converting {tag_to_convert}")
    converted_tag = wiegand.convert_to_wiegand(tag_to_convert)
    print(f"Converted: {converted_tag}")
    print(f"Converting back: {wiegand.convert_from_wiegand(converted_tag)}")


if __name__ == "__main__":
    main()
