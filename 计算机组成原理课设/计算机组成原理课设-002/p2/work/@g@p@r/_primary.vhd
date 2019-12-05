library verilog;
use verilog.vl_types.all;
entity GPR is
    port(
        inReg           : in     vl_logic_vector(31 downto 0);
        rs              : in     vl_logic_vector(4 downto 0);
        rt              : in     vl_logic_vector(4 downto 0);
        rd              : in     vl_logic_vector(4 downto 0);
        flag_Write      : in     vl_logic;
        clk             : in     vl_logic;
        rs1             : out    vl_logic_vector(31 downto 0);
        rt1             : out    vl_logic_vector(31 downto 0)
    );
end GPR;
