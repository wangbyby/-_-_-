library verilog;
use verilog.vl_types.all;
entity IM is
    port(
        address         : in     vl_logic_vector(31 downto 0);
        outIns          : out    vl_logic_vector(31 downto 0)
    );
end IM;
