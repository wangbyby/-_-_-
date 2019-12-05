library verilog;
use verilog.vl_types.all;
entity ALU is
    port(
        clk             : in     vl_logic;
        inA             : in     vl_logic_vector(31 downto 0);
        inB             : in     vl_logic_vector(31 downto 0);
        aluSelect       : in     vl_logic_vector(2 downto 0);
        outC            : out    vl_logic_vector(31 downto 0);
        flag_overflow_pos: out    vl_logic;
        flag_zero       : out    vl_logic
    );
end ALU;
