library verilog;
use verilog.vl_types.all;
entity Controll is
    port(
        op              : in     vl_logic_vector(5 downto 0);
        order_func      : in     vl_logic_vector(5 downto 0);
        sel_ALU         : out    vl_logic_vector(3 downto 0);
        beq             : out    vl_logic;
        Extop           : out    vl_logic;
        j               : out    vl_logic;
        jal             : out    vl_logic;
        jr              : out    vl_logic;
        lui             : out    vl_logic;
        rd              : out    vl_logic;
        imm_to_ALU      : out    vl_logic;
        GPR_write       : out    vl_logic;
        RAM_write       : out    vl_logic;
        RAM_to_GPR      : out    vl_logic
    );
end Controll;
