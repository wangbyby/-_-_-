module JExt( 
    input [25:0] target,input [31:0] pc, output [31:0] pcOut
);
assign pcOut = {pc[31:28],target,2'b00}; // jal j  
endmodule // JExtinput [25:0] target, input flag_sel, output [31:0] extTarget