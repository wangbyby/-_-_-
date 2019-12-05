module PC( 
    input clk,pcReset,
    input [31:0] pc ,
    output reg [31:0] pcOut);

always @(posedge clk or  pcReset) begin
    if (pcReset) begin
        pcOut <= 8'h0000_3000;
    end else begin
        pcOut <= pc; 
    end
end

endmodule // PC input clk