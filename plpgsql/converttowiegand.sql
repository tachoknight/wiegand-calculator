CREATE OR REPLACE FUNCTION "public"."converttowiegand" (p_num integer)  RETURNS integer
  VOLATILE
AS $body$
DECLARE
    v_baseVal VARCHAR(24) := '';
    v_fc VARCHAR(8) := '';
    v_uc VARCHAR(16) := '';

    v_fNum INTEGER;
    v_uNum INTEGER;

    v_FinalNum varchar(16) := '';
BEGIN
    -- Convert the number passed to us as a binary string
    v_baseVal := CAST(p_num::bit(24)::VARCHAR AS VARCHAR(24));
    -- Okay, we need two parts, the facility code, and the user code
    v_fc := SUBSTRING(v_baseVal from 1 for 8);
    v_uc := SUBSTRING(v_baseVal from 9);
    
    -- Now we're going to convert the bits to numbers
    v_fNum := (v_fc::bit(8))::integer;
    v_uNum := (v_uc::bit(16))::integer;
  
    -- And put it all together    
    v_FinalNum := format('%s%s', v_fNum::varchar, v_uNum::varchar);
  
    RETURN (SELECT v_FinalNum::integer);
END;
$body$ LANGUAGE plpgsql