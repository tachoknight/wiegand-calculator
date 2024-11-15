import math


def convert_to_wiegand(tag_sn):
    try:
        # Convert the tag to binary and make sure there are 24 bits
        bins = format(tag_sn, "024b")

        # Now split into the facility code and user code
        frontb = bins[:8]
        backb = bins[-16:]

        # Convert binary strings to integers
        f = int(frontb, 2)
        b = int(backb, 2)

        return int(f"{f}{b}")
    except ValueError as e:
        return str(e)


def convert_from_wiegand(board_tag):
    try:
        # We need to convert the board tag to a string
        board_tag = str(board_tag)

        # Split the string into facility code and user code
        facility_code = board_tag[:-5]
        user_code = board_tag[-5:]

        # Convert facility code to integer
        fc_num = int(facility_code)

        # Convert the facility code to binary and pad with zeros
        fc_bins = format(fc_num, "08b")

        # This is the number of bits we are working with, which we work down to 0
        bit_countdown = 24

        # Facility Code
        fc_bit_table = [0] * 8
        fc_idx = 0
        for char in fc_bins:
            if char == "1":
                fc_bit_table[fc_idx] = int(math.pow(2, bit_countdown - 1))
            else:
                fc_bit_table[fc_idx] = 0

            fc_idx += 1
            bit_countdown -= 1

        fc_sum = sum(fc_bit_table)

        # User Code
        uc_num = int(user_code)
        uc_bins = format(uc_num, "016b")

        uc_bit_table = [0] * 16
        uc_idx = 0
        for char in uc_bins:
            if char == "1":
                uc_bit_table[uc_idx] = int(math.pow(2, bit_countdown - 1))
            else:
                uc_bit_table[uc_idx] = 0

            uc_idx += 1
            bit_countdown -= 1

        uc_sum = sum(uc_bit_table)

        return int(str(fc_sum + uc_sum))
    except ValueError as e:
        print(f"Whoops, got {e}")
        return str(e)
