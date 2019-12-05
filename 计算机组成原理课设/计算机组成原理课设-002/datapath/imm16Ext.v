module imm16Ext(
    input [15:0] imm16,
    input flag_sel,
    output [31:0] imm32);
wire [31:0] b,c;
assign b = {{16{imm16[15]}},imm16};
assign c = {{16{1'b0}},imm16};

assign imm32 = (flag_sel)?b:c;

endmodule