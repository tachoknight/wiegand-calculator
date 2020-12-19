CREATE OR REPLACE FUNCTION "public"."convertfromwiegand" (p_num integer)  RETURNS integer
  VOLATILE
AS $body$
DECLARE
    v_baseVal varchar(8);
    v_facilityCode varchar(3);
    v_userCode varchar(5);

    v_bitCountdown integer := 24;

    -- All the facility variables we use
    v_facilityBits varchar(8);
    v_fbVal varchar(1);
    v_facilityBitTable varchar array[8]; 
    v_fcPos integer := 1;
    v_facilitySum integer := 0;

    -- And all the user variables
    v_userBits varchar(255);
    v_ubVal varchar(1);
    v_userBitTable varchar array[16];
    v_ucPos integer := 1;
    v_userSum integer := 0;

BEGIN
    v_baseVal := p_num::VARCHAR(8);

    -- We have to be careful about the facility code because it could be 
    -- three digits or less, while the user code will always be five
    -- digits
    v_facilityCode := substring(v_baseVal from 1 for length(v_baseVal) - 5);
    v_userCode := SUBSTRING(v_baseVal from length(v_baseVal) - 4);
    --raise notice '[%] - [%]', v_facilityCode, v_userCode;

    -- Okay, here we go with all our bit-twiddling logic....

    ----------------------------------------------------------------------
    -- Facility Code Logic
    ----------------------------------------------------------------------
    v_facilityBits := v_facilityCode::Integer::bit(8)::varchar;

    for pos in 1..8 loop
        v_fbVal := substring(v_facilityBits from pos for 1);
        if v_fbVal = '1' THEN
            v_facilityBitTable[v_fcPos] = pow(2, v_bitCountdown - 1)::integer::varchar;
        ELSE
            v_facilityBitTable[v_fcPos] = '0';
        end if;

        v_fcPos := v_fcPos + 1;
        v_bitCountdown := v_bitCountdown - 1;
    end loop; 

    for var in array_lower(v_facilityBitTable, 1)..array_upper(v_facilityBitTable, 1) loop
        --raise notice '--> [%]', v_facilityBitTable[var];
        v_facilitySum := v_facilitySum + v_facilityBitTable[var]::INTEGER;
    end loop;

    ----------------------------------------------------------------------
    -- User Code Logic
    ----------------------------------------------------------------------
    v_userBits := v_userCode::INTEGER::bit(16)::VARCHAR;

    for pos in 1..16 loop
        v_ubVal := substring(v_userBits from pos for 1);
        if v_ubVal = '1' THEN
            v_userBitTable[v_ucPos] = pow(2, v_bitCountdown - 1)::integer::varchar;
        ELSE
            v_userBitTable[v_ucPos] = '0';
        end if;

        v_ucPos := v_ucPos + 1;
        v_bitCountdown := v_bitCountdown - 1;
    end loop; 

    for var in array_lower(v_userBitTable, 1)..array_upper(v_userBitTable, 1) loop
        --raise notice '--> [%]', v_userBitTable[var];
        v_userSum := v_userSum + v_userBitTable[var]::INTEGER;
    end loop;

    return (select v_facilitySum + v_userSum);
end;
$body$ LANGUAGE plpgsql