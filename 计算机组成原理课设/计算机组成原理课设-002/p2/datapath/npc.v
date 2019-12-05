module NPC( 
    input [31:0] pc, input [25:0] instr_index,input [31:0] gpr_rs,input [15:0] target,
    input j,jal,jr,beq,flag_zero,
    output reg [31:0] npc , 
    output  reg [31:0] jal_out
);
//测试时 , 重点关照
wire [31:0] tmp;
assign tmp = pc +4; // 
wire [13:0] a;
assign a = {14{target[15]}};
wire [31:0] b ;
assign b = {a,target,2'b00};


//assign npc = (c)? tmp + b:tmp; //beq
//assign npc = (j)? {tmp[31:28],instr_index,2'b00} : tmp; // j
//assign npc = (jal)? {{tmp[31:28]},instr_index,2'b00} : tmp; //jal
//assign npc = (jr)?  gpr_rs : tmp;//jr
//assign jal_out = (jal)? pc+8 : 32'b0;
initial
begin
  jal_out = 32'b0;
end
always @(*)begin
  
  case( {beq,j,jal,jr})
    'b1000:// beq && zero
        begin
          if(flag_zero)
            npc = tmp + b;
          else
            npc = pc + 4;
        end
    'b0100: //j
        begin
          npc = {tmp[31:28],instr_index,2'b00};
        end
    'b0010: // jal
        begin
          npc = {{tmp[31:28]},instr_index,2'b00};
          jal_out = pc + 8;
        end
    'b0001: //jr
        begin
          npc = gpr_rs;
        end
      default:
          npc = tmp;
  endcase
  
end

endmodule
