// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: p/bi.proto

package com.singulariti.io.grpc;

public final class BiProto {
  private BiProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
  }
  static com.google.protobuf.Descriptors.Descriptor
    internal_static_p_BiLog_descriptor;
  static
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_p_BiLog_fieldAccessorTable;
  static com.google.protobuf.Descriptors.Descriptor
    internal_static_p_BiResult_descriptor;
  static
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_p_BiResult_fieldAccessorTable;
  static com.google.protobuf.Descriptors.Descriptor
    internal_static_p_PerformPathIssue_descriptor;
  static
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_p_PerformPathIssue_fieldAccessorTable;
  static com.google.protobuf.Descriptors.Descriptor
    internal_static_p_DeviceInfo_descriptor;
  static
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_p_DeviceInfo_fieldAccessorTable;
  static com.google.protobuf.Descriptors.Descriptor
    internal_static_p_CommandResult_descriptor;
  static
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_p_CommandResult_fieldAccessorTable;
  static com.google.protobuf.Descriptors.Descriptor
    internal_static_p_CommandExecutePerformance_descriptor;
  static
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_p_CommandExecutePerformance_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\np/bi.proto\022\001p\"S\n\005BiLog\022\023\n\013projectName\030" +
      "\001 \001(\t\022\022\n\nactionName\030\002 \001(\t\022\021\n\ttimestamp\030\003" +
      " \001(\003\022\016\n\006detail\030\004 \001(\014\"\'\n\010BiResult\022\016\n\006resu" +
      "lt\030\001 \001(\010\022\013\n\003msg\030\002 \001(\t\"D\n\020PerformPathIssu" +
      "e\022\020\n\010query_id\030\001 \001(\t\022\036\n\026can_perform_path_" +
      "count\030\002 \001(\005\"\204\001\n\nDeviceInfo\022\021\n\tdevice_id\030" +
      "\001 \001(\t\022\024\n\014manufacturer\030\002 \001(\t\022\r\n\005model\030\003 \001" +
      "(\t\022\023\n\013app_version\030\004 \001(\t\022\022\n\nni_version\030\005 " +
      "\001(\t\022\025\n\ropen_app_time\030\006 \001(\003\"\202\001\n\rCommandRe" +
      "sult\022\024\n\014command_name\030\001 \001(\t\022\020\n\010query_id\030\002",
      " \001(\t\022\031\n\021use_search_result\030\003 \001(\010\022\033\n\023user_" +
      "choose_trigger\030\004 \001(\010\022\021\n\tdate_time\030\005 \001(\003\"" +
      "\305\001\n\031CommandExecutePerformance\022\022\n\nspeak_t" +
      "ime\030\001 \001(\003\022!\n\031receive_voice_result_time\030\002" +
      " \001(\003\022\031\n\021send_command_time\030\003 \001(\003\022\033\n\023recei" +
      "ve_result_time\030\004 \001(\003\022\024\n\014command_text\030\005 \001" +
      "(\t\022\020\n\010query_id\030\006 \001(\t\022\021\n\tdevice_id\030\007 \001(\t2" +
      "T\n\005MisBi\022\035\n\002Bi\022\010.p.BiLog\032\013.p.BiResult\"\000\022" +
      ",\n\014BiDeviceInfo\022\r.p.DeviceInfo\032\013.p.BiRes" +
      "ult\"\000B$\n\027com.singulariti.io.grpcB\007BiProt",
      "oP\001b\006proto3"
    };
    com.google.protobuf.Descriptors.FileDescriptor.InternalDescriptorAssigner assigner =
        new com.google.protobuf.Descriptors.FileDescriptor.    InternalDescriptorAssigner() {
          public com.google.protobuf.ExtensionRegistry assignDescriptors(
              com.google.protobuf.Descriptors.FileDescriptor root) {
            descriptor = root;
            return null;
          }
        };
    com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        }, assigner);
    internal_static_p_BiLog_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_p_BiLog_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_p_BiLog_descriptor,
        new java.lang.String[] { "ProjectName", "ActionName", "Timestamp", "Detail", });
    internal_static_p_BiResult_descriptor =
      getDescriptor().getMessageTypes().get(1);
    internal_static_p_BiResult_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_p_BiResult_descriptor,
        new java.lang.String[] { "Result", "Msg", });
    internal_static_p_PerformPathIssue_descriptor =
      getDescriptor().getMessageTypes().get(2);
    internal_static_p_PerformPathIssue_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_p_PerformPathIssue_descriptor,
        new java.lang.String[] { "QueryId", "CanPerformPathCount", });
    internal_static_p_DeviceInfo_descriptor =
      getDescriptor().getMessageTypes().get(3);
    internal_static_p_DeviceInfo_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_p_DeviceInfo_descriptor,
        new java.lang.String[] { "DeviceId", "Manufacturer", "Model", "AppVersion", "NiVersion", "OpenAppTime", });
    internal_static_p_CommandResult_descriptor =
      getDescriptor().getMessageTypes().get(4);
    internal_static_p_CommandResult_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_p_CommandResult_descriptor,
        new java.lang.String[] { "CommandName", "QueryId", "UseSearchResult", "UserChooseTrigger", "DateTime", });
    internal_static_p_CommandExecutePerformance_descriptor =
      getDescriptor().getMessageTypes().get(5);
    internal_static_p_CommandExecutePerformance_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_p_CommandExecutePerformance_descriptor,
        new java.lang.String[] { "SpeakTime", "ReceiveVoiceResultTime", "SendCommandTime", "ReceiveResultTime", "CommandText", "QueryId", "DeviceId", });
  }

  // @@protoc_insertion_point(outer_class_scope)
}