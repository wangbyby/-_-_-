module PC( 
    input clk,pcReset,
    input [31:0] pc ,
    output reg [31:0] pcOut);

always @(posedge clk or  pcReset) begin
    if (pcReset) begin
        pcOut <= {{16{1'b0}},4'b0011,{12{1'b0}}};
    end else begin
        pcOut <= pc; 
    end
end

endmodule // PC input clk