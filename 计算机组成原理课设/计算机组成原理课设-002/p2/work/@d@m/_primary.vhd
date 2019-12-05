library verilog;
use verilog.vl_types.all;
entity DM is
    port(
        clk             : in     vl_logic;
        we              : in     vl_logic;
        read            : in     vl_logic;
        addr            : in     vl_logic_vector(11 downto 2);
        in_d            : in     vl_logic_vector(31 downto 0);
        out_d           : out    vl_logic_vector(31 downto 0)
    );
end DM;
